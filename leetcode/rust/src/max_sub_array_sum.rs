use std::collections::HashMap;

fn max_sub_array_sum(nums: Vec<i64>) -> i64 {
    let mut current_sum: i64 = nums[0];
    let mut current_max: i64 = nums[0];

    for (i, _) in nums.iter().enumerate() {
        if i == 0 {
            continue;
        }
        let temp = nums[i] + current_sum;

        if nums[i] > temp {
            current_sum = nums[i]
        } else {
            current_sum = temp
        }

        if current_sum > current_max {
            current_max = current_sum
        }
    }

    return current_max;
}

pub fn test_max_sub_array_sum() {
    let test_cases = HashMap::from([
        (vec![-2, 1, -3, 4, -1, 2, 1, -5, 4], 6),
        (vec![5, 4, -1, 7, 8], 23),
        (vec![1], 1),
    ]);

    for (key, val) in test_cases {
        let result = max_sub_array_sum(key);

        assert_eq!(result, val)
    }
}
