use std::fs;
use regex::Regex;

fn main() {
    let dump_file = fs::read_to_string("input.txt").expect("should be able to read file");
    let ops :Vec<&str> = dump_file.rsplit("\n").collect();
    let mut resp :i32 = 0;
    let re = Regex::new(r"mul\((\d+),(\d+)\)").unwrap();
    let dg = Regex::new(r"\d+").unwrap();
    for op in ops{
        // now do the parsing
        for cap in re.captures_iter(op){
            println!("{}",cap.get(0).unwrap().as_str());
            let digits: Vec<i32> = dg.find_iter(cap.get(0).unwrap().as_str()).filter_map(|digits| digits.as_str().parse().ok()).collect();
            resp += digits[0]*digits[1];
        }
    }
    println!("{}",resp)
}