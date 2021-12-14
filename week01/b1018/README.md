# **체스판 다시 칠하기**

출처 : https://www.acmicpc.net/problem/1018

<table>
<td>시간 제한</td><td>2초</td>
<td>메모리 제한</td><td>128MB</td>
</table>

## **문제**

지민이는 자신의 저택에서 MN개의 단위 정사각형으로 나누어져 있는 M×N 크기의 보드를 찾았다. 어떤 정사각형은 검은색으로 칠해져 있고, 나머지는 흰색으로 칠해져 있다. 지민이는 이 보드를 잘라서 8×8 크기의 체스판으로 만들려고 한다.

체스판은 검은색과 흰색이 번갈아서 칠해져 있어야 한다. 구체적으로, 각 칸이 검은색과 흰색 중 하나로 색칠되어 있고, 변을 공유하는 두 개의 사각형은 다른 색으로 칠해져 있어야 한다. 따라서 이 정의를 따르면 체스판을 색칠하는 경우는 두 가지뿐이다. 하나는 맨 왼쪽 위 칸이 흰색인 경우, 하나는 검은색인 경우이다.

보드가 체스판처럼 칠해져 있다는 보장이 없어서, 지민이는 8×8 크기의 체스판으로 잘라낸 후에 몇 개의 정사각형을 다시 칠해야겠다고 생각했다. 당연히 8*8 크기는 아무데서나 골라도 된다. 지민이가 다시 칠해야 하는 정사각형의 최소 개수를 구하는 프로그램을 작성하시오.


</br>

## 입력

첫째 줄에 N과 M이 주어진다. N과 M은 8보다 크거나 같고, 50보다 작거나 같은 자연수이다. 둘째 줄부터 N개의 줄에는 보드의 각 행의 상태가 주어진다. B는 검은색이며, W는 흰색이다.



</br>

## 출력

첫째 줄에 지민이가 다시 칠해야 하는 정사각형 개수의 최솟값을 출력한다.


</br>


<table>

<td valign="top" width="50%">

## 예제입력1

> 8 8</br>
WBWBWBWB</br>
BWBWBWBW</br>
WBWBWBWB</br>
BWBBBWBW</br>
WBWBWBWB</br>
BWBWBWBW</br>
WBWBWBWB</br>
BWBWBWBW</br>
</td>

<td>

</td>

<td valign="top" width="50%">

## 예제 출력1

> 1

</td>
</table>


</br>


<table>

<td valign="top" width="50%">

## 예제입력2

> 10 13</br>
BBBBBBBBWBWBW</br>
BBBBBBBBBWBWB</br>
BBBBBBBBWBWBW</br>
BBBBBBBBBWBWB</br>
BBBBBBBBWBWBW</br>
BBBBBBBBBWBWB</br>
BBBBBBBBWBWBW</br>
BBBBBBBBBWBWB</br>
WWWWWWWWWWBWB</br>
WWWWWWWWWWBWB</br>
</td>

<td>

</td>

<td valign="top" width="50%">

## 예제 출력2

> 12

</td>
</table>

</br>

<table>

<td valign="top" width="50%">

## 예제입력3

> 8 8</br>
BWBWBWBW</br>
WBWBWBWB</br>
BWBWBWBW</br>
WBWBWBWB</br>
BWBWBWBW</br>
WBWBWBWB</br>
BWBWBWBW</br>
WBWBWBWB</br>
</td>

<td>

</td>

<td valign="top" width="50%">

## 예제 출력3

> 0

</td>
</table>

</br>

<table>

<td valign="top" width="50%">

## 예제입력4

> 9 23</br>
BBBBBBBBBBBBBBBBBBBBBBB</br>
BBBBBBBBBBBBBBBBBBBBBBB</br>
BBBBBBBBBBBBBBBBBBBBBBB</br>
BBBBBBBBBBBBBBBBBBBBBBB</br>
BBBBBBBBBBBBBBBBBBBBBBB</br>
BBBBBBBBBBBBBBBBBBBBBBB</br>
BBBBBBBBBBBBBBBBBBBBBBB</br>
BBBBBBBBBBBBBBBBBBBBBBB</br>
BBBBBBBBBBBBBBBBBBBBBBW</br>
</td>

<td>

</td>

<td valign="top" width="50%">

## 예제 출력4

> 31

</td>
</table>

</br>

<table>

<td valign="top" width="50%">

## 예제입력5

> 10 10</br>
BBBBBBBBBB</br>
BBWBWBWBWB</br>
BWBWBWBWBB</br>
BBWBWBWBWB</br>
BWBWBWBWBB</br>
BBWBWBWBWB</br>
BWBWBWBWBB</br>
BBWBWBWBWB</br>
BWBWBWBWBB</br>
BBBBBBBBBB</br>
</td>

<td>

</td>

<td valign="top" width="50%">

## 예제 출력5

> 0

</td>
</table>

</br>

<table>

<td valign="top" width="50%">

## 예제입력6

> 8 8</br>
WBWBWBWB</br>
BWBWBWBW</br>
WBWBWBWB</br>
BWBBBWBW</br>
WBWBWBWB</br>
BWBWBWBW</br>
WBWBWWWB</br>
BWBWBWBW</br>
</td>

<td>

</td>

<td valign="top" width="50%">

## 예제 출력6

> 2

</td>
</table>

</br>

<table>

<td valign="top" width="50%">

## 예제입력7

> 11 12</br>
BWWBWWBWWBWW</br>
BWWBWBBWWBWW</br>
WBWWBWBBWWBW</br>
BWWBWBBWWBWW</br>
WBWWBWBBWWBW</br>
BWWBWBBWWBWW</br>
WBWWBWBBWWBW</br>
BWWBWBWWWBWW</br>
WBWWBWBBWWBW</br>
BWWBWBBWWBWW</br>
WBWWBWBBWWBW</br>
</td>

<td>

</td>

<td valign="top" width="50%">

## 예제 출력7

> 15

</td>
</table>

</br>