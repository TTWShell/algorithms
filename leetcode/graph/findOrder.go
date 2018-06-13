/* https://leetcode.com/problems/course-schedule-ii/description/
There are a total of n courses you have to take, labeled from 0 to n - 1.

Some courses may have prerequisites, for example to take course 0 you have to first take course 1, which is expressed as a pair: [0,1]

Given the total number of courses and a list of prerequisite pairs, return the ordering of courses you should take to finish all courses.

There may be multiple correct orders, you just need to return one of them. If it is impossible to finish all courses, return an empty array.

For example:

    2, [[1,0]]

There are a total of 2 courses to take. To take course 1 you should have finished course 0. So the correct course order is [0,1]

    4, [[1,0],[2,0],[3,1],[3,2]]

There are a total of 4 courses to take. To take course 3 you should have finished both courses 1 and 2. Both courses 1 and 2 should be taken after you finished course 0. So one correct course order is [0,1,2,3]. Another correct ordering is[0,2,1,3].

Note:
    The input prerequisites is a graph represented by a list of edges, not adjacency matrices. Read more about how a graph is represented.
    You may assume that there are no duplicate edges in the input prerequisites.

Hints:
    This problem is equivalent to finding the topological order in a directed graph. If a cycle exists, no topological ordering exists and therefore it will be impossible to take all courses.
    Topological Sort via DFS - A great video tutorial (21 minutes) on Coursera explaining the basic concepts of Topological Sort.
    Topological sort could also be done via BFS.
*/

package lgraph

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	degree := make([]int, numCourses)

	for _, prereqs := range prerequisites {
		to, from := prereqs[0], prereqs[1]
		graph[from] = append(graph[from], to)
		degree[to]++
	}

	res, queue := make([]int, 0, numCourses), []int{}
	for i := range degree {
		if degree[i] == 0 {
			res, queue = append(res, i), append(queue, i)
		}
	}

	for len(queue) != 0 {
		tmpQueue := []int{}
		for _, from := range queue {
			for _, to := range graph[from] {
				degree[to]--
				if degree[to] == 0 {
					res, tmpQueue = append(res, to), append(tmpQueue, to)
				}
			}
		}
		queue = tmpQueue
	}

	if len(res) == numCourses {
		return res
	}
	return []int{}
}
