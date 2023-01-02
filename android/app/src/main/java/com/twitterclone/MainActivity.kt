package com.twitterclone

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.compose.material.Scaffold
import com.twitterclone.screens.chooseUsername.ChooseUsernameView
import com.twitterclone.screens.createAccount.CreateAccountScreenView
import com.twitterclone.screens.enterRegistrationPassword.EnterRegistrationPasswordScreenView
import com.twitterclone.screens.registrationPrompt.RegistrationPromptScreenView
import com.twitterclone.screens.uploadProfilePicture.UploadProfilePictureView
import com.twitterclone.screens.verifyEmailForRegistration.VerifyEmailForRegistrationView
import com.twitterclone.theme.AppTheme

class MainActivity : ComponentActivity( ) {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)

        setContent {
            AppTheme {

                UploadProfilePictureView( )

            }
        }
    }
}