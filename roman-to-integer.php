<?php
class Solution {

    /**
     * @param String $s
     * @return Integer
     */
    public static $roman = [
            'I'	=>	1,
            'V'	=>	5,
            'X'	=>	10,
            'L'	=>	50,
            'C'	=>	100,
            'D'	=>	500,
            'M'	=>	1000
        ];
    
    function romanToInt($s) {
        
       $strlen = strlen($s);
        $num= $prev_num = static::$roman[$s[$strlen - 1]];
        for($i = $strlen -2; $i >=0 ; $i--){
            $current = static::$roman[$s[$i]];
            if ($current < $prev_num) {
                $num -= $current;
            }else{
                $num += $current;
            }
            $prev_num = $current;
        }
        return $num;
    }

}


?>  
