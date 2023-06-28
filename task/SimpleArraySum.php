<?php

    function SimpleArraySum(array $data):int
    {
        $result = 0;
        for ($i=0; $i < count($data) ; $i++) { 
            $result += $data[$i];
        }

        return $result;
    }