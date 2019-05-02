extern crate unicode_segmentation;

use unicode_segmentation::UnicodeSegmentation;

pub fn reverse(input: &str) -> String {
    let g = UnicodeSegmentation::graphemes(input, true)
        .rev()
        .collect::<Vec<&str>>();

    return g.join("");
}
