import random


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

    def __init__(self, suit_index, num_index):
        self.suit = card.possible_suits[suit_index]
        self.num = card.possible_nums[num_index]


def create_deck():
    cards = []
    for suit_index in range(len(card.possible_suits)):
        for num_index in range(len(card.possible_nums)):
            cards.append(card(suit_index, num_index))
    return cards


def shuffle_deck(cards):
    random.shuffle(cards)
    return cards


def print_deck(cards):
    for card in cards:
        print(f"{card.num} of {card.suit}")


class gamestate:
    def __init__(self):
        self.my_hand = []
        self.cards_revealed = []
        self.my_money = 1000
        self.opponent_money = 1000


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
    # Returns a value from 0 to 1 based on how good my hand is compared to other possible hands
    return 1


def main():
    deck_of_cards = shuffle_deck(create_deck())
    print_deck(deck_of_cards)


main()
