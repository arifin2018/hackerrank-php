<?php

// Path to run ./vendor/bin/phpunit --bootstrap vendor/autoload.php FileName.php
// Butuh Framework PHPUnit

// ./vendor/bin/phpunit --bootstrap vendor/autoload.php TestAVeryBigSum.php

use PHPUnit\Framework\TestCase;

// Class yang mau di TEST.
require_once "task/AVeryBigSum.php";

// Class untuk run Testing.
class TestAVeryBigSum extends TestCase
{
    public function testCountWords()
    {
        $WordCount = AVeryBigSum([1000000001,1000000002,1000000003,1000000004,1000000005]);

        $this->assertEquals(5000000015, $WordCount); 
    }
}