<?php

// Path to run ./vendor/bin/phpunit --bootstrap vendor/autoload.php FileName.php
// Butuh Framework PHPUnit

// ./vendor/bin/phpunit --bootstrap vendor/autoload.php Solvemefirst.php

use PHPUnit\Framework\TestCase;

// Class yang mau di TEST.
require_once "task/PlusMinus.php";

// Class untuk run Testing.
class TestPlusMinus extends TestCase
{
    public function testCountWords()
    {
        $WordCount = PlusMinus([-4, 3, -9, 0, 4]);

        $this->expectOutputString(0.500000, $WordCount); 
    }
}