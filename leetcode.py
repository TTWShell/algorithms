import argparse
from collections import namedtuple
import datetime
import json
import sqlite3
import time

import requests
from tqdm import tqdm

conn = sqlite3.connect("leetcode.db")
cur = conn.cursor()
table_name = 'problem'
cur.execute(f"""
CREATE TABLE IF NOT EXISTS `{table_name}` (
    id integer(5) NOT NULL PRIMARY KEY,
    title varchar(128),
    title_slug varchar(128),
    difficulty tinyint(1),
    status varchar(16) DEFAULT NULL,
    go tinyint(1),
    updated_at DATETIME DEFAULT (
        strftime('%Y-%m-%d %H:%M:%S', 'now', 'localtime'))
    )""")
cur.execute(f"""
CREATE TRIGGER IF NOT EXISTS [UpdateLastTime]
    AFTER
    UPDATE
    ON `{table_name}`
    FOR EACH ROW
    WHEN NEW.updated_at <= OLD.updated_at
BEGIN
    update `{table_name}` set updated_at=strftime(
    '%Y-%m-%d %H:%M:%S', 'now', 'localtime') where id=OLD.id;
END""")
Prombem = namedtuple('Problem', ['id', 'title', 'title_slug', 'difficulty', 'status', 'go', 'updated_at'])


class LeetCode:

    def __init__(self, username, password, difficulty, index):
        self.username = username
        self.password = password
        self.difficulty = difficulty
        self.index = index
        self.session = requests.Session()
        self.login()

    def _get_csrfmiddlewaretoken(self):
        url = 'https://leetcode.com/graphql'
        headers = {
            'Referer': 'https://leetcode.com/playground/UpwhGDg6/shared',
            'x-csrftoken': '39EmD1AhCNr2GcIffQQ5v5FF45l3tJ59IoBkNuXIFXKEqERelZnEx61xoGQo0zhR',
        }
        cookies = {
            'csrftoken': '39EmD1AhCNr2GcIffQQ5v5FF45l3tJ59IoBkNuXIFXKEqERelZnEx61xoGQo0zhR',
        }
        payload = {
            "operationName": "globalData",
            "variables": {},
            "query": "query globalData {\n userStatus {\nisSignedIn\nusername\n}\n}\n"
        }
        resp = self.session.post(url, payload, headers=headers, cookies=cookies)
        assert resp.status_code == 200, f'excepted status_code 200, {resp.status_code} occurred: {resp.text}'
        return resp.cookies['csrftoken']

    def _request(self, url):
        for i in range(3):
            resp = self.session.get(url)
            if resp.status_code == 429:  # Too Many Requests
                time.sleep(30)
                print('>>> 429 Too Many Requests: {} retry...'.format(i))
                continue
            return resp

    def login(self):
        url = 'https://leetcode.com/accounts/login/'
        headers = {
            'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8',
            'Accept-Encoding': 'gzip, deflate, sdch',
            'Accept-Language': 'en-US,en;q=0.8,zh-CN;q=0.6,zh;q=0.4,zh-TW;q=0.2',
            'Connection': 'keep-alive',
            'Host': 'leetcode.com',
            'Referer': 'https://leetcode.com/accounts/login/',
            'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.11; rv:56.0) Gecko/20100101 Firefox/56.0',
            'Content-Type': 'application/x-www-form-urlencoded',
        }
        data = {
            'csrfmiddlewaretoken': self._get_csrfmiddlewaretoken(),
            'login': self.username,
            'password': self.password,
        }
        resp = self.session.post(url, data=data, headers=headers)
        assert resp.status_code == 200, f'login faield'

    def profile(self):
        """profile data
        """
        url = 'https://leetcode.com/api/progress/all/'
        resp = self.session.get(url)
        assert resp.status_code == 200, f'request profile failed: {resp.status_code}'

        data = resp.json()
        print('{unsolved} Todo | {solvedTotal}/{questionTotal} Solved'
              ' | {attempted} Attempted'.format(**data))
        print('Easy {Easy} Medium {Medium} Hard {Hard}'.format(
            **data['solvedPerDifficulty']))

    def support_languages_by_tilte_slug(self, tilte_slug):
        desc_url = 'https://leetcode.com/problems/{}/description/'.format(
            tilte_slug)
        resp = self._request(desc_url)
        csrftoken = resp.cookies['csrftoken']

        url = 'https://leetcode.com/graphql'
        headers = {
            'Accept': '*/*',
            'Accept-Encoding': 'gzip',  # 'gzip, deflate, br',
            'Accept-Language': 'zh-CN,zh;q=0.9,en;q=0.8',
            'Connection': 'keep-alive',
            'Referer': desc_url,
            'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.11; rv:56.0) Gecko/20100101 Firefox/56.0', # NOQA E501
            'Content-Type': 'application/json',
            'x-csrftoken': csrftoken
        }
        payload = {
            "query": '\n'.join([
                'query getQuestionDetail($titleSlug: String!) {',
                '  isCurrentUserAuthenticated',
                '  question(titleSlug: $titleSlug) {',
                '    codeDefinition',
                '  }',
                '}',
            ]),
            "variables": {
                "titleSlug": tilte_slug,
            },
            "operationName": "getQuestionDetail"
        }
        resp = self.session.post(url, json=payload, headers=headers)
        assert resp.status_code == 200, f'query codeDefinition failed: {resp.status_code}'

        code_definition = resp.json()['data']['question']['codeDefinition']
        return [code['text'] for code in json.loads(code_definition)]

    def sync_all_problems(self):
        url = 'https://leetcode.com/api/problems/algorithms/'
        data = self.session.get(url).json()
        today = datetime.datetime.now()

        retrieve_sql = 'SELECT * FROM `{table_name}` WHERE id={pk}'
        insert_sql = 'INSERT INTO `{table_name}` (id, title, title_slug, difficulty, status, go) VALUES ' + \
            '({id}, "{title}", "{title_slug}", {difficulty}, "{status}", {go})'
        update_sql = 'UPDATE `{table_name}` SET go={go}, status="{status}" WHERE id={pk}'

        pairs = tqdm(sorted(data['stat_status_pairs'], key=lambda x: x['stat']['frontend_question_id']))
        pairs = filter(lambda x: x['paid_only'] is False, pairs)
        if self.difficulty:
            pairs = filter(lambda x: x['difficulty']['level'] == self.difficulty, pairs)

        for p in pairs:
            # question_id && frontend_question_id
            problem_id = p['stat']['frontend_question_id']
            title_slug = p['stat']['question__title_slug']
            status = p['status']

            problem = cur.execute(retrieve_sql.format(table_name=table_name, pk=problem_id)).fetchone()
            if problem is None:
                go = 1 if 'Go' in self.support_languages_by_tilte_slug(title_slug) else 0
                values = {
                    'id': problem_id,
                    'title': p['stat']['question__title'],
                    'title_slug': title_slug,
                    'difficulty': p['difficulty']['level'],
                    'status': status,
                    'go': go,
                }
                cur.execute(insert_sql.format(table_name=table_name, **values))
                conn.commit()
                continue

            problem = Prombem(*problem)
            need_update = False
            if status is not None and problem.status != status:
                print(f'Latest {status} problem: {problem.id}. {problem.title}')
                need_update = True

            go = problem.go
            if go != 1 and (today - datetime.datetime.strptime(problem.updated_at, '%Y-%m-%d %H:%M:%S')).days >= 7:
                need_update = True
                if 'Go' in self.support_languages_by_tilte_slug(title_slug):
                    go = 1

            if need_update:
                cur.execute(update_sql.format(table_name=table_name, go=go, pk=problem_id, status=status))
                conn.commit()

    def get_unsolved_problem(self):
        """
        Get a problem by difficulty filter.
        """
        self.sync_all_problems()
        all_sql = 'SELECT * FROM `{table_name}` ORDER BY id'.format(table_name=table_name)
        pairs = []
        for p in cur.execute(all_sql).fetchall():
            problem = Prombem(*p)
            if problem.status != 'ac' and problem.go and (
                    self.difficulty is None or
                    problem.difficulty == self.difficulty):
                pairs.append(problem._asdict())

        print('UnSolved:', len(pairs))
        if pairs:
            index = self.index % len(pairs)
            print(pairs[index])
            print(f'https://leetcode.com/problems/{pairs[index]["title_slug"]}/description/')

    def analysis(self, profile, query):
        if profile:
            self.profile()
        if query:
            self.get_unsolved_problem()


parser = argparse.ArgumentParser(
    description='Analysis leetcode problems progress.')
parser.add_argument('-u', '--username', type=str, required=True,
                    help='Username for login.')
parser.add_argument('-p', '--password', type=str, required=True,
                    help='Password to use when connecting to leetcode.')
parser.add_argument('-d', '--difficulty', type=int,
                    choices=(1, 2, 3),
                    help='Difficulty of problems. 1-3: easy --> hard')
parser.add_argument('-i', '--index', type=int, default=0,
                    help='Choose an unsolved problem, default is first.')

parser.add_argument('-P', '--profile', action='store_true',
                    help='Show profile, all statistical information.')
parser.add_argument('-query', '--query', action='store_true',
                    help='Query one unsolved problem order by id.')


if __name__ == '__main__':
    args = parser.parse_args()
    leetcode = LeetCode(username=args.username, password=args.password, difficulty=args.difficulty, index=args.index)
    leetcode.analysis(profile=args.profile, query=args.query)
