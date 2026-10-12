use std::collections::HashMap;

fn remove_outer_parenthesis(str: String) -> String {
    let mut stack: Vec<char> = Vec::new();
    let mut result: String = String::new();

    for val in str.chars() {
        if val == '(' {
            stack.push(val);

            if stack.len() > 1 {
                result += "(";
            }
        } else {
            stack.pop();

            if stack.len() > 0 {
                result += ")";
            }
        }
    }

    return result;
}

pub fn test_remove_outer_parenthesis() {
    let test_cases = HashMap::from([
        ("(()())(())", "()()()"),
        ("(()())(())(()(()))", "()()()()(())"),
        ("()()", ""),
    ]);

    for (key, value) in test_cases {
        let result = remove_outer_parenthesis(key.to_string());

        assert_eq!(result, value);
    }
}
