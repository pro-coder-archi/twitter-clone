package com.twitterclone.navigation

import androidx.compose.runtime.Composable
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.compose.rememberNavController
import com.twitterclone.screens.createAccount.CreateAccountScreenView
import com.twitterclone.screens.registrationPrompt.RegistrationPromptScreenView

@Composable
fun AppNavigation( ) {
    val navController= rememberNavController( )

    NavHost(
        navController = navController,
        startDestination = Screen.RegistrationPromptScreen.route
    ) {

        composable(route = Screen.RegistrationPromptScreen.route) {
            RegistrationPromptScreenView( )
        }

        composable(route = Screen.CreateAccountScreen.route) {
            CreateAccountScreenView( )
        }

    }
}