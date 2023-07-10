<?php

// Path to run ./vendor/bin/phpunit --bootstrap vendor/autoload.php FileName.php
// Butuh Framework PHPUnit

// ./vendor/bin/phpunit --bootstrap vendor/autoload.php TestBreakingTheRecords.php

use PHPUnit\Framework\TestCase;

// Class yang mau di TEST.
require_once "task/BreakingTheRecords.php";

// Class untuk run Testing.
class TestBreakingTheRecords extends TestCase
{
    public function testBreakingRecords()
    {
        $data = breakingRecords([10,5,20,20,4,5,2,25,1]);
        $data2 = breakingRecords([3,4,21,36,10,28,35,5,24,42]);

        $this->assertEquals([2,4], $data); 
        $this->assertEquals([4,0], $data2); 
    }
}