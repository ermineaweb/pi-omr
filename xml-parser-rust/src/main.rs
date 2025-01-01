use std::fs;

fn main() {
    let data = fs::read_to_string("carmen.xml").expect("Unable to read file");
    // println!("{}", data);

    let opt = roxmltree::ParsingOptions {
        allow_dtd: true,
        ..roxmltree::ParsingOptions::default()
    };

    // Find element by id.
    let doc = roxmltree::Document::parse_with_options(&data, opt).expect("error");

    println!("{:?}", doc);

    let elem = doc
        .descendants()
        .find(|n| n.attribute("id") == Some("rect1"))
        .expect("error");

    println!("{:?}", elem);

    assert!(elem.has_tag_name("rect"));
}
