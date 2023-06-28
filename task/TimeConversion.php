<?php

function TimeConversion($s){
    // $regex = "/^(1[0-2]|0?[1-9]):[0-5][0-9]:[0-5][0-9](AM)$/i";

    // if (preg_match($regex, $s)) {
    //     echo $regex;

    // }else{
    //     echo 'hallo';
    // }
    return date("H:i:s", strtotime($s));
}

echo TimeConversion("07:05:45PM");