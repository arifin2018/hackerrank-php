<?php

function breakingRecords(array $scores):array {
    $highestScore = $scores[0];
    $lowestScore = $scores[0];
    $countHighestScore = 0;
    $countLowestScore = 0;

    foreach ($scores as $key => $value) {
        if ($value > $highestScore) {
            $countHighestScore++;
            $highestScore = $scores[$key];
        }elseif ($value < $lowestScore) {
            $countLowestScore++;
            $lowestScore = $scores[$key];
        }
    }

    return [$countHighestScore, $countLowestScore];
}

breakingRecords([10,5,20,20,4,5,2,25,1]);