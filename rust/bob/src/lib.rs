const QUESTION_ANSWER: &str = "Sure.";
const EMPTY_ANSWER: &str = "Fine. Be that way!";
const YELL_QUESTION_ANSWER: &str = "Calm down, I know what I'm doing!";
const DEFAULT_ANSWER: &str = "Whatever.";
const YELL_ANSWER: &str = "Whoa, chill out!";

pub fn reply(message: &str) -> &str {
    let sanitized_message: String = message.trim().replace("\n", "").replace("\t", "");
    let is_uppercase = sanitized_message.to_uppercase().eq(&sanitized_message)
        && !sanitized_message
            .to_uppercase()
            .eq(&sanitized_message.to_lowercase());
    let is_question = sanitized_message.trim().ends_with("?");

    if is_question && is_uppercase {
        return YELL_QUESTION_ANSWER;
    } else if is_uppercase {
        return YELL_ANSWER;
    } else if is_question {
        return QUESTION_ANSWER;
    } else if sanitized_message.trim().chars().count() == 0 {
        return EMPTY_ANSWER;
    }
    return DEFAULT_ANSWER;
}
