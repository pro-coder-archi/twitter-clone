package com.twitterclone.screens.registrationPrompt

import androidx.compose.foundation.Image
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.*
import androidx.compose.material.Divider
import androidx.compose.material.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.unit.dp
import com.twitterclone.theme.AppTheme
import com.twitterclone.R
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.sp
import com.twitterclone.components.button.ButtonComponentVariants
import com.twitterclone.components.button.ButtonComponentView

@Composable
fun RegistrationPromptScreenView( ) {

    Column(modifier = Modifier.fillMaxSize( ).padding(horizontal = 16.dp, vertical = 10.dp)) {

        Box(
            modifier = Modifier.fillMaxWidth( ).weight(1f),
            contentAlignment = Alignment.Center
        ) {
            Image(
                painter = painterResource(id = R.drawable.twitter_logo),
                contentDescription = ""
            )
        }

        Text(
            text = "See what is happening in the world right now !",
            style = AppTheme.typography.blackHeading,
            modifier = Modifier.widthIn(0.dp, 350.dp),
        )

        Spacer(modifier = Modifier.height(100.dp))

        ButtonComponentView(
            buttonText = "Continue with Google",
            clickHandler = { },
            leadingDrawableResourceID = R.drawable.google_logo,
            variant = ButtonComponentVariants.WIDE
        )

        Row(
            verticalAlignment = Alignment.CenterVertically,
            modifier = Modifier.fillMaxWidth( ).padding(vertical= 5.dp),
            horizontalArrangement = Arrangement.Center
        ) {
            Divider(modifier = Modifier.width(100.dp).background(AppTheme.colors.text), thickness = 1.dp)

            Text(
                text = "or",
                modifier = Modifier.padding(horizontal = 10.dp),
                style = AppTheme.typography.smallSizedBoldText.copy(fontSize = 14.sp, color = AppTheme.colors.text)
            )

            Divider(modifier = Modifier.width(100.dp).background(AppTheme.colors.text), thickness = 1.dp)
        }

        ButtonComponentView(
            buttonText = "Create new account",
            clickHandler = { },
            variant = ButtonComponentVariants.WIDE
        )

        Spacer(modifier = Modifier.height(20.dp))

        Text(

            text = buildAnnotatedString
                {
                    withStyle(style = SpanStyle(color = AppTheme.colors.text))
                        {append("By signing up, you agree to our ")}

                    withStyle(style = SpanStyle(color = AppTheme.colors.primary))
                        {append("Terms, Privacy Policy ")}

                    withStyle(style = SpanStyle(color = AppTheme.colors.text))
                        {append("and")}

                    withStyle(style = SpanStyle(color = AppTheme.colors.primary))
                        {append(" Cookie Use")}
                },

            style = AppTheme.typography.smallSizedParagraphText
        )

        Spacer(modifier = Modifier.height(35.dp))

        Text(

            text = buildAnnotatedString
                {
                    withStyle(style = SpanStyle(color = AppTheme.colors.text))
                        {append("Have an account already? ")}

                    withStyle(style = SpanStyle(color = AppTheme.colors.primary))
                        {append("Sign in")}
                },

            style = AppTheme.typography.smallSizedParagraphText
        )
    }
}