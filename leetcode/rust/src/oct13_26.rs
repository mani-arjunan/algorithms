fn max_alternating_sum(nums: Vec<i64>) -> i64 {
    let mut result: i64 = 0;
    let mut new_nums: Vec<i64> = Vec::new();

    for val in nums {
        new_nums.push(val * val)
    }

    new_nums.sort_by(|a, b| b.cmp(a));
    let mut mid_point = new_nums.len() / 2;

    if new_nums.len() % 2 == 0 {
        mid_point = mid_point - 1;
    }

    let mut i = 0;

    while i <= mid_point {
        result += new_nums[i];
        i += 1;
    }

    let mut j = mid_point + 1;
    while j < new_nums.len() {
        result -= new_nums[j];
        j += 1;
    }

    return result;
}

pub fn test_max_alternating_sum() {
    assert_eq!(12, max_alternating_sum(vec![1, 2, 3]));
    assert_eq!(16, max_alternating_sum(vec![1, -1, 2, -2, 3, -3]));
}
