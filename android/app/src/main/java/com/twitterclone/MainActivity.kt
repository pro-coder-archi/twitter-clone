package com.twitterclone

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import com.twitterclone.screens.registrationPrompt.RegistrationPromptScreenView
import com.twitterclone.theme.AppTheme

class MainActivity : ComponentActivity( ) {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)

        setContent {
            AppTheme {

                RegistrationPromptScreenView( )

            }
        }
    }
}