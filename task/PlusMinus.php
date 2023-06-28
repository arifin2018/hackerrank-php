<?php

function PlusMinus(array $arr)
{
    $resultPositif = 0;
    $resultNegatif = 0;
    $resultZero = 0;

    foreach ($arr as $key => $value) {
        if ($value == 0) {
            $resultZero += 1;
        }elseif ($value < 0) {
            $resultNegatif += 1;
        }else{
            $resultPositif += 1;
        }
    }

    $resultPositif /= count($arr);
    $resultNegatif /= count($arr);
    $resultZero /= count($arr);

    echo sprintf('%8f',$resultPositif) . PHP_EOL;
    echo sprintf('%8f',$resultNegatif). PHP_EOL;
    echo sprintf('%8f',$resultZero). PHP_EOL;
}

PlusMinus([-4, 3, -9, 0, 4, 1]);