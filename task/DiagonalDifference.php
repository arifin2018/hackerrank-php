<?php

function diagonalDifference(array $arr):int
{
    $key1 = 0;
    $key2 = count($arr)-1;
    $data1 = 0;
    $data2 = 0;
    $result = 0;
    for ($i=0; $i < count($arr) ; $i++) {
        $data1 += $arr[$i][$key1];
        $data2 += $arr[$i][$key2];
        $key1++;
        $key2--;
    }
    $result = $data1 - $data2;
    return abs($result);
}