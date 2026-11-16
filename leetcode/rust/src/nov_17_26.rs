fn digit_sum(str: String, k: i32) -> String {
    let mut copy_s = str;

    while copy_s.len() > k as usize {
        let times = (copy_s.len() as i32 + (k - 1)) / k;
        let mut temp: String = String::new();
        let mut i = 0;

        while i < times {
            let start = i * k;
            let mut str = String::from("");

            if times - 1 == i {
                if let Some(t) = copy_s.get((start as usize)..copy_s.len()) {
                    str = t.to_string();
                }
            } else {
                if let Some(t) = copy_s.get((start as usize)..(start + k) as usize) {
                    str = t.to_string();
                }
            }

            let mut add_result = 0;

            for i in str.chars() {
                add_result += i.to_digit(10).unwrap();
            }

            temp += &add_result.to_string();
            i += 1;
        }

        copy_s = temp;
    }

    return copy_s;
}

pub fn test_digit_sum() {
    assert_eq!(String::from("000"), digit_sum(String::from("00000000"), 3));
    assert_eq!(
        String::from("135"),
        digit_sum(String::from("11111222223"), 3)
    );
}
