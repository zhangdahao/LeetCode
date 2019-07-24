<?php
class Solution {

    /**
     * @param Integer $x
     * @return Boolean
     */
    function isPalindrome($x){
        $left = 1;
        $right = strlen($x);
        while ($left < $right) {
            if (substr($x, -$left, 1) !== substr($x, -$right, 1)) {
                return false;
            }
            $left++;
            $right--;
        }
        return true;
    }
}
?>
