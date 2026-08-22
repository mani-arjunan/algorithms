use std::{cell::RefCell, rc::Rc};

struct Node {
    value: i16,
    next: Option<Rc<RefCell<Node>>>,
}

struct Queue {
    length: i16,
    head: Option<Rc<RefCell<Node>>>,
    tail: Option<Rc<RefCell<Node>>>,
}

impl Queue {
    fn enqueue(&mut self, value: i16) {
        let node = Rc::new(RefCell::new(Node {
            value: value,
            next: None,
        }));

        self.length += 1;

        match &mut self.tail {
            Some(t) => {
                t.borrow_mut().next = Some(node.clone());
                *t = node.clone();
            }
            None => {
                self.head = Some(node.clone());
                self.tail = Some(node.clone());
                return;
            }
        }
    }

    fn dequeue(&mut self) -> i16 {
        if self.head.is_none() {
            return -1;
        }

        self.length -= 1;

        if let Some(t) = &self.tail {
            if let Some(h) = &self.head {
                if Rc::ptr_eq(&h, &t) {
                    let value = h.borrow().value;
                    self.head = None;
                    self.tail = None;

                    return value;
                }

                let next = h.borrow().next.clone();
                let value = h.borrow().value;

                self.head = next;
                return value;
            }
        }

        return -1;
    }
}

pub fn test_queue() {
    let mut queue = Queue {
        length: 0,
        head: None,
        tail: None,
    };

    queue.enqueue(1);
    queue.enqueue(2);
    queue.enqueue(3);
    println!("Dequeue Value: {}", queue.dequeue());

    match &queue.head {
        Some(h) => {
            println!("{}", h.borrow().value);
            if let Some(h1) = &h.borrow().next {
                let next = h1.borrow().next.clone();
                println!("{}", h1.borrow().value);
                match next {
                    None => {
                        println!("Queue Ended")
                    }
                    Some(_) => {}
                }
            }
        }
        None => {}
    }
}
