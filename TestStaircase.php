<?php

// Path to run ./vendor/bin/phpunit --bootstrap vendor/autoload.php FileName.php
// Butuh Framework PHPUnit

// ./vendor/bin/phpunit --bootstrap vendor/autoload.php Solvemefirst.php

use PHPUnit\Framework\TestCase;

// Class yang mau di TEST.
require_once "task/Staircase.php";

// Class untuk run Testing.
class TestStaircase extends TestCase
{
    public function testCountWords()
    {
        $WordCount = Staircase(6);

        $this->expectOutputString(0.500000, $WordCount); 
    }
}