<?php

function birthdayCakeCandles(array $number)
{
    rsort($number);
    if (count($number) == 0) {
        return 0;
    }
    $result = 0;
    $maxData = $number[0];
    foreach ($number as $key => $value) {
        if ($maxData < $value || $maxData == $value) {
            $maxData = $value;
            $result++;
        }
    }
    return $result;
}

echo birthdayCakeCandles([18,90,90,13,90,75,90,8,90,43]);