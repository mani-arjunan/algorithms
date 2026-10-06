use std::collections::HashMap;

fn min_addition_to_make_valid_string(str: String) -> i16 {
    let mut stack: Vec<char> = Vec::new();
    let mut close_stack: Vec<char> = Vec::new();

    for char in str.chars() {
        if char == '(' {
            stack.push('(');
        } else {
            if stack.len() == 0 {
                close_stack.push(')');
            } else {
                stack.pop();
            }
        }
    }

    let result: i16 = (stack.len() + close_stack.len()) as i16;

    return result;
}

pub fn test_oct06_26() {
    let test_cases = HashMap::from([
        ("())", 1),
        ("(((", 3),
        (")))", 3),
        (")())(())", 2),
        ("(())((", 2),
    ]);

    for (key, value) in test_cases {
        let result = min_addition_to_make_valid_string(key.to_string());

        assert_eq!(result, value);
    }
}
