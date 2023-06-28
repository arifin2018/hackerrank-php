<?php

// Path to run ./vendor/bin/phpunit --bootstrap vendor/autoload.php FileName.php
// Butuh Framework PHPUnit

// ./vendor/bin/phpunit --bootstrap vendor/autoload.php TestMiniMaxSum.php

use PHPUnit\Framework\TestCase;

// Class yang mau di TEST.
require_once "task/miniMaxSum.php";

// Class untuk run Testing.
class TestMiniMaxSum extends TestCase
{
    public function testCountWords()
    {
        // $WordCount1 = miniMaxSum([1,2,3,4,5]);
        $WordCount2 = miniMaxSum([396285104,573261094,759641832,819230764,364801279]);

        // $this->expectOutputString("10 14", $WordCount1); 
        $this->expectOutputString("2093989309 2548418794", $WordCount2); 
    }
}