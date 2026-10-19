fn target_indices(mut nums: Vec<i16>, target: i16) -> Vec<usize> {
    let mut result: Vec<usize> = Vec::new();
    nums.sort_by(|a, b| a.cmp(b));

    for (i, val) in nums.iter().enumerate() {
        if *val == target {
            result.push(i);
        }
    }

    return result;
}

pub fn test_target_indices() {
    assert_eq!(vec![1, 2], target_indices(vec![1, 2, 5, 2, 3], 2));
    assert_eq!(vec![3], target_indices(vec![1, 2, 5, 2, 3], 3));
    assert_eq!(vec![4], target_indices(vec![1, 2, 5, 2, 3], 5));
}
