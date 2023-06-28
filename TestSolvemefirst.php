<?php

// Path to run ./vendor/bin/phpunit --bootstrap vendor/autoload.php FileName.php
// Butuh Framework PHPUnit

// ./vendor/bin/phpunit --bootstrap vendor/autoload.php Solvemefirst.php

use PHPUnit\Framework\TestCase;

// Class yang mau di TEST.
require_once "task/Solvemefirst.php";

// Class untuk run Testing.
class TestSolvemefirst extends TestCase
{
    public function testCountWords()
    {
        $WordCount = Solvemefirst(4,4);

        $this->assertEquals(8, $WordCount); 
        $this->assertEquals(4, $WordCount); 
    }
}