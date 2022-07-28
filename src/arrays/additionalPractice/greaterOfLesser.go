package additionalPractice

/*
  Problem : https://www.interviewbit.com/problems/greater-of-lesser/

  Solution :

  1) General Approach

  -  Iterate through each input arrays, with given condition
  -  return max of both conditions

*/

// T(n) : O(n); S(n) : O(1)
func countGreaterOrLesser(A []int, B []int, C int) int {

	nA, nB, cA, cB := len(A), len(B), 0, 0
	for i := 0; i < nA; i++ {
		if A[i] > C {
			cA++
		}
	}
	for i := 0; i < nB; i++ {
		if B[i] < C {
			cB++
		}
	}

	if cA > cB {
		return cA
	}

	return cB
}

func SolveCountGreaterOrLesser(A []int, B []int, C int) int {
	return countGreaterOrLesser(A, B, C)
}