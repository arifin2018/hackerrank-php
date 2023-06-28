<?php

function getTotalX(array $a,array $b)
{
    $lowIntB = 0;
    
    if ($b != null) {
        $lowIntB = $b[0];
        foreach ($b as $key => $value) {
            if ($lowIntB >= $value) {
                $lowIntB = $value;
            }
        }
    }
    $dataA = arrayA($a,$lowIntB);
    $dataB = arrayB($b,$dataA);
    return count($dataB);
}

function arrayA($a,$lowIntB) {
    $dataA = [];
    $resultDataA = [];
    
    foreach ($a as $key => $value) {
        for ($i=1; $i <= $lowIntB; $i++) { 
            if ($i % $value == 0) {
                array_push($dataA, $i); 
            }
        }
    }
    if (count($a) < 2) {
        return $dataA;
    }

    for ($i=0; $i < count($dataA); $i++) { 
        for ($j=$i+1; $j < count($dataA); $j++) { 
            if ($dataA[$i] == $dataA[$j]) {
                array_push($resultDataA, $dataA[$i]); 
            }
        }
    }
    if (count($a) < 3) {
        if (isset($a[0])) {
            return $resultDataA;
        }
    }

    $obj= [];
    $resultData = [];
    for ($i=0; $i < count($resultDataA) ; $i++) { 
        for ($j=$i+1; $j <count($resultDataA); $j++) { 
            if ($resultDataA[$i] == $resultDataA[$j]) {
                if (!isset($obj[$resultDataA[$i]])) {
                    $obj[$resultDataA[$i]] = 0;
                }
                $obj[$resultDataA[$i]] = $obj[$resultDataA[$i]] + 1;
            }
        }
    }
    foreach ($obj as $key => $value) {
        array_push($resultData, $key);
    }
    return $resultData;
}

function arrayB($b,$dataA) {
    $checkingZero = true;
    $dataB = [];
    $countDatab = [];
    foreach ($b as $keyB => $valueB) {
        foreach ($dataA as $keyA => $valueA) {
            if ($valueB % $valueA != 0) {
                $checkingZero = false;
            }else{
                array_push($dataB, $valueA);
            }
        }
        $checkingZero = true;   
    }
    if (count($b) < 2) {
        return $dataB;
    }

    for ($i=0; $i <count($dataB); $i++) { 
        for ($j=$i+1; $j < count($dataB) ; $j++) { 
            if ($dataB[$i] == $dataB[$j]) {
                if (!isset($countDatab[$dataB[$i]])) {
                    $countDatab[$dataB[$i]] = 0;
                }
                $countDatab[$dataB[$i]]++;
            }
        }
    }
    $result = [];
    $resultCheckingZero = 0;

    foreach ($countDatab as $key => $value) {
        if ($value > $resultCheckingZero || $value == $resultCheckingZero) {
            $resultCheckingZero = $value;
            array_push($result, $key);
        }
    }
    return $result;
}


print_r(getTotalX([2,4],[16,32,96]));
