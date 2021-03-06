# **최소 대기 시간**

## **문제**

비어있지 않은 시간의 길이를 의미하는 자연수로 이루어진 특정 배열이 있다. 순서에 상관없이 한 번에 하나의 쿼리만 실행 가능하다.
</br>
시간의 길이로 정의되어 있는 쿼리의 대기시간은 반드시 그 전 쿼리가 실행할 때까지 기다려야 한다.
</br>
최소 대기시간을 갖는 값을 반환하여라.
</br>
예를 들어, [1, 4, 5]의 배열이 주어지고, 쿼리의 실행순서가 [5, 1, 4]이라면, 최종 대기시간은 (0) + (5) + (5 + 1) = 11이다. 첫 번째 쿼리는 즉시 실행되므로 0, 두 번째 쿼리 대기시간은 첫 번째 쿼리의 대기시간인 5만큼 대기하고, 마지막 쿼리의 대기시간은 그 전 모든 쿼리 대기시간의 합인 6이다.

> 주의 : input 배열의 순서를 바꿀 수 있다.

</br>

## 예제 입력

```go
queries = [3, 2, 1, 2, 6]
```

## 예제 출력

```go
17
```