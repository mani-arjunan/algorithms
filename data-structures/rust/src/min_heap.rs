struct Heap {
    data: Vec<i32>,
    length: usize,
}

impl Heap {
    fn insert(&mut self, value: i32) {
        self.data.push(value);
        self.heapify_up(self.length);
        self.length += 1;
    }

    fn delete(&mut self) -> i32 {
        if self.length == 0 {
            return -1;
        }

        let head_element = self.data[0];

        if self.length == 1 {
            self.data = Vec::new();
            self.length = 0;
            return head_element;
        }

        self.data[0] = self.data[self.data.len() - 1];
        self.length -= 1;
        self.heapify_down(0);
        self.data.pop();

        return head_element;
    }

    fn get_left_child_index(&mut self, index: usize) -> usize {
        return (2 * index) + 1;
    }

    fn get_right_child_index(&mut self, index: usize) -> usize {
        return (2 * index) + 2;
    }

    fn get_parent(&mut self, index: usize) -> usize {
        return (index - 1) / 2;
    }

    fn heapify_up(&mut self, index: usize) {
        if index == 0 {
            return;
        }

        let parent_index = self.get_parent(index);
        let parent_value = self.data[parent_index];
        let current_value = self.data[index];

        if current_value < parent_value {
            self.data[parent_index] = current_value;
            self.data[index] = parent_value;
            self.heapify_up(parent_index);
        }
    }

    fn heapify_down(&mut self, index: usize) {
        let left_child_index = self.get_left_child_index(index);
        let right_child_index = self.get_right_child_index(index);

        if index >= self.length || left_child_index >= self.length {
            return;
        }

        let left_value = self.data[left_child_index];
        let current_value = self.data[index];

        if right_child_index >= self.length {
            if current_value > left_value {
                self.data[index] = left_value;
                self.data[left_child_index] = current_value;
            }
            return;
        }

        let right_value = self.data[right_child_index];

        if right_value <= left_value && current_value > right_value {
            self.data[index] = right_value;
            self.data[right_child_index] = current_value;
            self.heapify_down(right_child_index);
        } else if right_value > left_value && current_value > left_value {
            self.data[index] = left_value;
            self.data[left_child_index] = current_value;
            self.heapify_down(left_child_index);
        }
    }
}

pub fn test_cost_minimization() {
    let arr = vec![5, 3, 5, 2];
    let mut heap = Heap {
        data: Vec::new(),
        length: 0,
    };
    let mut total = 0;

    for n in arr {
        heap.insert(n);
    }

    while heap.length > 1 {
        let elem1 = heap.delete();
        let elem2 = heap.delete();

        total += elem1 + elem2;
        heap.insert(elem1 + elem2);
    }

    println!("Final Value => {}", total);
}
