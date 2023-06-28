<?php

// Path to run ./vendor/bin/phpunit --bootstrap vendor/autoload.php FileName.php
// Butuh Framework PHPUnit

// ./vendor/bin/phpunit --bootstrap vendor/autoload.php SimpleArraySum.php

use PHPUnit\Framework\TestCase;

// Class yang mau di TEST.
require_once "task/ComparetheTriplets.php";

// Class untuk run Testing.
class TestComparetheTriplets extends TestCase
{
    public function testCountWords()
    {
        $WordCount1 = ComparetheTriplets([17,28,30], [99,16,8]);
        $WordCount2 = ComparetheTriplets([5,6,7], [3,6,10]);

        $this->assertEquals([2,1], $WordCount1); 
        $this->assertEquals([1,1], $WordCount2); 
    }
}