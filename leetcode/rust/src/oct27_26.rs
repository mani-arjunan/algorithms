fn create_grid(mut m: i16, n: i16) -> Vec<String> {
    let mut result: Vec<String> = Vec::new();

    while m > 0 {
        let mut str: String = String::new();
        let mut temp = n;

        while temp > 0 {
            if m == 1 {
                str.push('.');
            } else {
                if temp == n {
                    str.push('.')
                } else {
                    str.push('#')
                }
            }
            temp -= 1;
        }
        result.push(str);
        m -= 1;
    }

    return result;
}

pub fn test_create_grid() {
    assert_eq!(vec!["...."], create_grid(1, 4));
    assert_eq!(vec![".##", "..."], create_grid(2, 3));
    assert_eq!(vec![".##", ".##", "..."], create_grid(3, 3));
}
