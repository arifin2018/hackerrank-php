<?php

function AVeryBigSum(array $arr):int
{
    $result = 0;

    foreach ($arr as $key => $value) {
        $result += $value;
    }

    return $result;
}
