class Solution {

    /**
     * @param String[] $strs
     * @return String
     */
    function longestCommonPrefix($strs) {
        if (empty($strs) || !is_array($strs)) {
            return "";
        }

        $flag = 1;
        $str_len = strlen($strs[0]);
        while ($flag <= $str_len) {
            $fix = substr($strs[0], 0, $flag);
            foreach ($strs as $str) {
                if(strpos($str, $fix) !== 0){
                    return substr($str, 0, $flag - 1);
                }
            }
            $flag++;
        }
        return $strs[0];
    }
}
