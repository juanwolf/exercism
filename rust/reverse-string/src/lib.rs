extern crate unicode_segmentation;

use unicode_segmentation::UnicodeSegmentation;

pub fn reverse(input: &str) -> String {
    let mut g: Vec<&str> = UnicodeSegmentation::graphemes(input, true).collect::<Vec<&str>>();
    g.reverse();

    return g.join("");
}
