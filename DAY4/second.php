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
            // XMAS as an X
            if ($mat[$i][$j] == "A"){
                if (((($i - 1) >= 0) && ($j - 1) >= 0) && (($i + 1 < $width) && ($j +1) < $width)){
                    $majDiag = [$mat[$i-1][$j-1],$mat[$i][$j],$mat[$i+1][$j+1]];
                    $minDag  = [$mat[$i+1][$j-1],$mat[$i][$j],$mat[$i-1][$j+1]];
                    if (checkMAS($majDiag) && checkMAS($minDag)){
                        $resp++;
                    }
                }
            }
        }
    }
    echo $resp;

    function checkMAS($diag){
        $mas = ["M","A","S"];
        $sam = ["S","A","M"];
        if (($mas === $diag) || ($sam === $diag)){
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
