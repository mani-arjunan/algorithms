fn sum_of_squares(nums: Vec<i16>) -> i16 {
    let n = nums.len();
    let mut total = 0;

    for (index, val) in nums.iter().enumerate() {
        if n % (index + 1) == 0 {
            total += val * val;
        }
    }

    return total;
}

pub fn test_oct12_26() {
    assert_eq!(21, sum_of_squares(vec![1, 2, 3, 4]));
    assert_eq!(63, sum_of_squares(vec![2, 7, 1, 19, 18, 3]));
}
