<?php

function ComparetheTriplets(array $a, array $b):array
{
    $result = [0,0];

    for ($i=0; $i < count($a) ; $i++) { 
        if ($a[$i] == $b[$i]) {
            false;
        }elseif ($a[$i] > $b[$i]) {
            $result[0] += 1;
        } 
        else{
            $result[1] += 1;
        }
    }
    return $result;
}

