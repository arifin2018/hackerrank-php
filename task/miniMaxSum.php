<?php

function miniMaxSum(array $arr):void
{
    sort($arr);
    $min = 0;
    $max = 0;

    if (!empty($arr)) {
        $datamin = $arr;
        unset($datamin[0]);
        foreach ($datamin as $key => $value) {
            $min += $value;
        }
        $datamax = $arr;
        unset($datamax[count($datamax)-1]);
        foreach ($datamax as $key => $value) {
            $max += $value;
        }
    }else{
        false;
    }
    
    echo "$max $min";
}

miniMaxSum([256741038,623958417,467905213,714532089,938071625]);