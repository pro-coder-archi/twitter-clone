package com.twitterclone.navigation

sealed class Screen(val route: String) {

    object RegistrationPromptScreen: Screen("/registration-prompt")
    object CreateAccountScreen: Screen("/create-account")
}