class card:
    possible_suits = ["Hearts", "Diamonds", "Spades", "Clubs"]
    possible_nums = [
        "Ace",
        "Two",
        "Three",
        "Four",
        "Five",
        "Six",
        "Seven",
        "Eight",
        "Nine",
        "Ten",
        "Jack",
        "Queen",
        "King",
    ]  # Aces both high and low
    suit = possible_suits[0]
    num = possible_nums[0]

    def create_card(suit_index, num_index):
        card.suit = card.possible_suits[suit_index]
        card.num = card.possible_nums[num_index]
        return card


class deck:
    cards = []

    def create_deck():
        for suit_index in range(len(card.possible_suits)):
            for num_index in range(len(card.possible_nums)):
                deck.cards.append(card.create_card(suit_index, num_index))
        return deck.cards

    def print_deck(cards):
        for card in cards:
            print(f"{card.num} of {card.suit}")


class gamestate:
    my_hand = []
    cards_revealed = []


def what_do_next(gamestate):
    # Returns call/match bet, check, bet + amount, or fold
    pussy = False
    next_do = "check"
    if pussy:
        next_do = "fold"
    elif gamestate.previous_bet > gamestate.my_curr_bet:
        next_do = "call"
    return next_do


def get_hand_rating(gamestate):
    return 1


def main():
    deck_of_cards = deck.create_deck()
    deck.print_deck(deck_of_cards)


main()
