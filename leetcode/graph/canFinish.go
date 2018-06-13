/* https://leetcode.com/problems/course-schedule/description/
There are a total of n courses you have to take, labeled from 0 to n - 1.

Some courses may have prerequisites, for example to take course 0 you have to first take course 1, which is expressed as a pair: [0,1]

Given the total number of courses and a list of prerequisite pairs, is it possible for you to finish all courses?

For example:

    2, [[1,0]]
There are a total of 2 courses to take. To take course 1 you should have finished course 0. So it is possible.

    2, [[1,0],[0,1]]
There are a total of 2 courses to take. To take course 1 you should have finished course 0, and to take course 0 you should also have finished course 1. So it is impossible.

Note:
    The input prerequisites is a graph represented by a list of edges, not adjacency matrices. Read more about how a graph is represented.
    You may assume that there are no duplicate edges in the input prerequisites.
*/

package lgraph

func canFinish(numCourses int, prereqs [][]int) bool {
	graph := make([][]int, numCourses)
	for i := 0; i < len(prereqs); i++ {
		graph[prereqs[i][1]] = append(graph[prereqs[i][1]], prereqs[i][0])
	}

	visited := make([]int, numCourses)
	for course := range graph {
		if !canFinishHelper(course, graph, visited) {
			return false
		}
	}
	return true
}

func canFinishHelper(course int, graph [][]int, visited []int) bool {
	const (
		partial  = 1
		complete = 2
	)
	if visited[course] == partial {
		return false
	} else if visited[course] == complete {
		return true
	}
	visited[course] = partial
	for _, dependentCourse := range graph[course] {
		if !canFinishHelper(dependentCourse, graph, visited) {
			return false
		}
	}
	visited[course] = complete
	return true
}
