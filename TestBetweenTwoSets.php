<?php

// Path to run ./vendor/bin/phpunit --bootstrap vendor/autoload.php TestBetweenTwoSets.php
// Butuh Framework PHPUnit

// ./vendor/bin/phpunit --bootstrap vendor/autoload.php TestBetweenTwoSets.php

use PHPUnit\Framework\TestCase;

// Class yang mau di TEST.
require_once "task/BetweenTwoSets.php";

// Class untuk run Testing.
class TestBetweenTwoSets extends TestCase
{
    public function testgetTotalX()
    {
        $getTotalX = getTotalX([1],[100]);
        $getTotalX2 = getTotalX([3,9,6],[36,72]);
        $getTotalX3 = getTotalX([2,4],[16,32,96]);

        $this->assertEquals(9, $getTotalX); 
        $this->assertEquals(2, $getTotalX2); 
        $this->assertEquals(3, $getTotalX3); 
    }
}