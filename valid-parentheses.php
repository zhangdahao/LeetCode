<?php
class Solution {
    /**
     * @param String $s
     * @return Boolean
     */
     function isValid($s){
        $strlen = strlen($s);
        if ($strlen & 1) return false;

        $positive = [
            '('	=>	1,
            '[' =>  2,
            '{' =>  3
        ];

        $negative = [
            ')'	=>	1,
            ']'	=>	2,
            '}'	=>	3
        ];

        $stack = [];
        for($i = 0; $i<$strlen; $i++){
            if (isset($positive[$s[$i]])) {
                $stack[] = $positive[$s[$i]];
            }else{
                if(array_pop($stack) !== $negative[$s[$i]]){
                    return false;
                }
            }
        }

        if (empty($stack)) {
            return true;
        }
        return false;
    }
}
?>
