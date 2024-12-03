<?php
    $file = fopen("input.txt","r") or die("unable to read input");
    $content = fread($file,filesize("input.txt"));
    fclose($file);    
    $mat = loadIntoMatrix($content);
    $width = sizeof($mat[1]);
    $height = sizeof($mat);
    $resp = 0;
    for ($i = 0; $i < sizeof($mat[1]); $i++) {
        for ($j = 0; $j < sizeof($mat); $j++) {
            //$mat[$i][$j];
            // $i has rows
            // $j has columns
            // vertical downwards
            if (($i + 3) < $height){
                if (checkXmas($mat[$i][$j],$mat[$i+1][$j],$mat[$i+2][$j],$mat[$i+3][$j])){
                    $resp++;
                }
            }
            // check horizontal right
            if (($j + 3) < $width){
                if (checkXmas($mat[$i][$j],$mat[$i][$j+1],$mat[$i][$j+2],$mat[$i][$j+3])){
                    $resp++;
                }
            }
            // check vertical upwards
            if (($i - 3) >= 0){
                if (checkXmas($mat[$i][$j],$mat[$i-1][$j],$mat[$i-2][$j],$mat[$i-3][$j])){
                    $resp++;
                }
            }
            // check horizontal left 
            if (($j - 3) >= 0){
                if (checkXmas($mat[$i][$j],$mat[$i][$j-1],$mat[$i][$j-2],$mat[$i][$j-3])){
                    $resp++;
                }
            }
            // check NE
            if (($j + 3) < $width && ($i - 3) >= 0){
                if (checkXmas($mat[$i][$j],$mat[$i-1][$j+1],$mat[$i-2][$j+2],$mat[$i-3][$j+3])){
                    $resp++;
                }
            }
            // check SE
            if (($j + 3) < $width && ($i + 3) < $height){
                if (checkXmas($mat[$i][$j],$mat[$i+1][$j+1],$mat[$i+2][$j+2],$mat[$i+3][$j+3])){
                    $resp++;
                }
            }
            // check SW
            if (($j - 3) >= 0 && ($i + 3) < $height){
                if (checkXmas($mat[$i][$j],$mat[$i+1][$j-1],$mat[$i+2][$j-2],$mat[$i+3][$j-3])){
                    $resp++;
                }
            }
            // check NW
            if (($j - 3) >= 0 && ($i - 3) >= 0){
                if (checkXmas($mat[$i][$j],$mat[$i-1][$j-1],$mat[$i-2][$j-2],$mat[$i-3][$j-3])){
                    $resp++;
                }
            }
        }
    }
    echo $resp;

    function checkXmas($a,$b,$c,$d){
        if (($a == "X") && ($b == "M") && ($c == "A") && ($d == "S")){
            return true;
        }
        return false;
    }

    function loadIntoMatrix($content){
        // split in rows
        $height = 0;
        $matrix = [];
        $rows = explode("\n",$content);
        // as safe as an unexploded bomb
        foreach ($rows as $row){
            $j = 0;
            foreach (str_split($row) as $c){
                $matrix[$height][$j] = $c;
                $j++;
            }
            $height++;
        }
        return $matrix;
    }
?>