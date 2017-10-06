#! /usr/local/bin/python
import argparse
import os
import pprint

import requests

pp = pprint.PrettyPrinter()


class LeetCode:
    LOGIN_URI = 'https://leetcode.com/accounts/login/'
    HEADERS = {
        'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8',  # NOQA E501
        'Accept-Encoding': 'gzip, deflate, sdch',
        'Accept-Language': 'en-US,en;q=0.8,zh-CN;q=0.6,zh;q=0.4,zh-TW;q=0.2',
        'Connection': 'keep-alive',
        'Host': 'leetcode.com',
        'Referer': 'https://leetcode.com/accounts/login/',
        'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.11; rv:56.0) Gecko/20100101 Firefox/56.0', # NOQA E501
        'Content-Type': 'application/x-www-form-urlencoded',
    }

    def __init__(self, username, password, difficulty):
        self.username = username
        self.password = password
        self.difficulty = difficulty
        self.session = requests.Session()
        self.login()

    def _get_csrfmiddlewaretoken(self):
        resp = self.session.get(self.LOGIN_URI, headers=self.HEADERS)
        # <input type='hidden' name='csrfmiddlewaretoken' value='B5yVhk7grKZU4rt3hNxGZ9NUB67lmIW0RuAg75Vx9f3WA8grGIV4qI2eFbmlK6Dq' />  # NOQA E501
        return resp.text.split('csrfmiddlewaretoken')[1].split("'")[2]

    def login(self):
        data = {
            'csrfmiddlewaretoken': self._get_csrfmiddlewaretoken(),
            'login': self.username,
            'password': self.password,
        }
        resp = self.session.post(
            self.LOGIN_URI, data=data, headers=self.HEADERS)
        if resp.status_code != 200:
            print("Login failed.", resp.status_code)
            os._exit(1)

    def all(self):
        url = 'https://leetcode.com/api/progress/all/'
        resp = self.session.get(url)
        if resp.status_code != 200:
            print("request failed:", resp.status_code)
            os._exit(1)
        data = resp.json()
        print('{unsolved} Todo | {solvedTotal}/{questionTotal} Solved'
              ' | {attempted} Attempted'.format(**data))
        print('Easy {Easy} Medium {Medium} Hard {Hard}'.format(
            **data['solvedPerDifficulty']))

    def get_unsolved_problem(self):
        # get a problem by difficulty filter
        level_maps = {
            "easy": 1,
            "medium": 2,
            "hard": 3,
        }

        url = 'https://leetcode.com/api/problems/all/'
        data = self.session.get(url).json()

        pairs = sorted(data['stat_status_pairs'],
                       key=lambda x: x['stat']['question_id'])
        pairs = filter(lambda x: x['paid_only'] is False, pairs)
        pairs = filter(lambda x: x['status'] is None, pairs)
        if self.difficulty:
            level = level_maps[self.difficulty]
            pairs = filter(lambda x: x['difficulty']['level'] == level, pairs)

        pp.pprint(pairs[0])
        print('https://leetcode.com/problems/{}/description/'.format(
            pairs[0]['stat']['question__title_slug']))

    def analysis(self, allstatistics, query):
        if allstatistics:
            self.all()
        if query:
            self.get_unsolved_problem()


parser = argparse.ArgumentParser(
    description='Analysis leetcode problems progress.')
parser.add_argument('-u', '--username', type=str, required=True,
                    help='Username for login.')
parser.add_argument('-p', '--password', type=str, required=True,
                    help='Password to use when connecting to leetcode.')
parser.add_argument('-d', '--difficulty', type=str,
                    choices=('easy', 'medium', 'hard'),
                    help='Difficulty of problems.')

parser.add_argument('-a', '--allstatistics', action='store_true',
                    help='Show all statistical information.')
parser.add_argument('-query', '--query', action='store_true',
                    help='Query one unsolved problem order by id.')


if __name__ == '__main__':
    args = parser.parse_args()
    leetcode = LeetCode(username=args.username, password=args.password,
                        difficulty=args.difficulty)
    leetcode.analysis(allstatistics=args.allstatistics, query=args.query)
