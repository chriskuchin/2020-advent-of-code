use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;
use regex::Regex;

fn main() {
    // File hosts must exist in current path before this produces output
    let re = Regex::new(r"(?P<min>\d+)-(?P<max>\d+) (?P<val>\w)").unwrap();
    if let Ok(lines) = read_lines("./input") {
        // Consumes the iterator, returns an (Optional) String
        let mut valid_passwords = 0;
        for line in lines {
            if let Ok(pw) = line {
                let rule: Vec<&str> = pw.rsplit(":").collect();
                println!("{}", rule[1]);

                let cap = re.captures(rule[1]).unwrap();
                let min_str = cap.name("min").map_or("", |min| min.as_str());
                let max_str = cap.name("max").map_or("", |max| max.as_str());
                let val = cap.name("val").map_or("", |val| val.as_str());

                let min = &min_str.parse::<i32>().unwrap_or(0);
                let max = &max_str.parse::<i32>().unwrap_or(0);

                // let mut occurrences = 0;
                // for c in rule[0].chars() {
                //     if c == char::from(val.as_bytes()[0]) {
                //         occurrences += 1;
                //     }
                // }

                let min_char = rule[0].chars().nth(*min as usize).unwrap();
                let max_char = rule[0].chars().nth(*max as usize).unwrap();
                let val_char = char::from(val.as_bytes()[0]);

                if min_char != max_char && (min_char == val_char || max_char == val_char) {
                    valid_passwords += 1;
                }
                // if occurrences >= *min && occurrences <= *max {
                //     valid_passwords += 1;
                // }
            }
        }
        println!("{}", valid_passwords)
    }
}

// The output is wrapped in a Result to allow matching on errors
// Returns an Iterator to the Reader of the lines of the file.
fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}