<?php

// Path to run ./vendor/bin/phpunit --bootstrap vendor/autoload.php FileName.php
// Butuh Framework PHPUnit

// ./vendor/bin/phpunit --bootstrap vendor/autoload.php TestDiagonalDifference.php

use PHPUnit\Framework\TestCase;

// Class yang mau di TEST.
require_once "task/DiagonalDifference.php";

// Class untuk run Testing.
class TestDiagonalDifference extends TestCase
{
    public function testCountWords()
    {
        $WordCount = diagonalDifference([[11,2,4],[4,5,6],[10,8,-12]]);

        $this->assertEquals(15, $WordCount); 
    }
}