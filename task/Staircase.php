<?php

function Staircase(int $n):void
{
    $count2 = $n;
    for ($i=1; $i <= $n; $i++) { 
        for ($j=0; $j<$count2-1; $j++) { 
            echo ' ';
        }
        for ($k=0; $k < $i; $k++) { 
            echo '#';
        }
        echo PHP_EOL;
        $count2--;
    }
}

Staircase(6);
/*
0
01
012
*/
