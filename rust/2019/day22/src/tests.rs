mod tests {
    use crate::part1;
    use crate::part2;
    use day22::Card;

    #[test]
    fn test_part1_1() {
        let mut card = Card::new(0, 10);
        card.deal(7);
        card.new_stack();
        card.new_stack();
        assert_eq!(0, card.get_pos());
        let mut card = Card::new(7, 10);
        card.deal(7);
        card.new_stack();
        card.new_stack();
        assert_eq!(9, card.get_pos());
    }

    #[test]
    fn test_part1_2() {
        let mut card = Card::new(3, 10);
        card.cut(6);
        card.deal(7);
        card.new_stack();
        assert_eq!(0, card.get_pos());
        let mut card = Card::new(6, 10);
        card.cut(6);
        card.deal(7);
        card.new_stack();
        assert_eq!(9, card.get_pos());
    }

    #[test]
    fn test_part1_3() {
        let mut card = Card::new(9, 10);
        card.new_stack();
        card.cut(-2);
        card.deal(7);
        card.cut(8);
        card.cut(-4);
        card.deal(7);
        card.cut(3);
        card.deal(9);
        card.deal(3);
        card.cut(-1);
        assert_eq!(0, card.get_pos());
        let mut card = Card::new(6, 10);
        card.new_stack();
        card.cut(-2);
        card.deal(7);
        card.cut(8);
        card.cut(-4);
        card.deal(7);
        card.cut(3);
        card.deal(9);
        card.deal(3);
        card.cut(-1);
        assert_eq!(9, card.get_pos());
    }

    #[test]
    fn test_input_part1() {
        assert_eq!(3749, part1(include_str!("../input.txt").to_string()));
    }

    #[test]
    fn test_input_part2() {
        assert_eq!(
            77225522112241,
            part2(include_str!("../input.txt").to_string())
        );
    }
}
