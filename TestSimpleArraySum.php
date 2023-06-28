<?php

// Path to run ./vendor/bin/phpunit --bootstrap vendor/autoload.php FileName.php
// Butuh Framework PHPUnit

// ./vendor/bin/phpunit --bootstrap vendor/autoload.php SimpleArraySum.php

use PHPUnit\Framework\TestCase;

// Class yang mau di TEST.
require_once "task/SimpleArraySum.php";

// Class untuk run Testing.
class TestSimpleArraySum extends TestCase
{
    public function testCountWords()
    {
        $WordCount1 = AVeryBigSum([1000000001,1000000002,1000000003,1000000004,1000000005]);

        $this->assertEquals(31, $WordCount1); 
    }
}