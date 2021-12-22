# **Rest Stops**

출처 : https://www.acmicpc.net/problem/15748

<table>
<td>시간 제한</td><td>2초</td>
<td>메모리 제한</td><td>512MB</td>
</table>

## **문제**

전체길이 L 미터 (1<=L<=10^6)
John: 속도 Rf seconds/meter
Bessie: 속도 Rb seconds/meter
Rest stops : N개, xi (위치), ci(해당위치에서 얻을 수 있는 tastiness unit)

특정 i번째 rest stop에 t초 동안 머물렀다면 Ci * t 만큼 tatiness 값을 얻을 수 있다.

John은 Bessie 보다 느리다
Bessie는 John보다 뒤에 있을 수 없다
Bessie가 얻을 수 있는 최대 총 tastiness units는?

</br>

## 입력

첫 줄에 L, N, Rf, Rb가 주어지고, 다음 줄에 각 rest stops의 위치(xi)와 얻을 수 있는 tastiness unit(ci)가 주어진다.

</br>

## 출력

A single integer: the maximum total tastiness units Bessie can obtain.


</br>


## 예제입력1

```
10 2 4 3
7 2
8 1
```

## 예제 출력1
```
15
```
