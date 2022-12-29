package com.twitterclone.screens.registrationPrompt

import androidx.compose.foundation.Image
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.Button
import androidx.compose.material.ButtonDefaults
import androidx.compose.material.Divider
import androidx.compose.material.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.unit.dp
import com.twitterclone.theme.AppTheme
import com.twitterclone.R
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.TextStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.sp

@Composable
fun RegistrationPromptScreenView( ) {

    Column(modifier = Modifier.fillMaxSize( ).padding(16.dp)) {

        Box(
            modifier = Modifier.fillMaxWidth( ).weight(1f),
            contentAlignment = Alignment.Center
        ) {
            Image(
                painter = painterResource(id = R.drawable.twitter_logo),
                contentDescription = "Twitter logo"
            )
        }

        Text(
            text = "See what is happening in the world right now !",
            style = AppTheme.typography.blackHeading,
            modifier = Modifier.widthIn(0.dp, 350.dp)
        )

        Spacer(modifier = Modifier.height(100.dp))

        Button(
            onClick = { },
            shape = RoundedCornerShape(35.dp),
            colors = ButtonDefaults.buttonColors(backgroundColor = Color.White),
            modifier = Modifier.fillMaxWidth( ).height(42.5.dp)
        ) {
            Image(
                painter = painterResource(id = R.drawable.google_logo),
                contentDescription = "Google logo",
                modifier = Modifier.height(height = 25.dp)
            )

            Spacer(modifier = Modifier.width(width = 10.dp))

            Text(
                text = "Continue with Google",
                color = Color.Black,
                lineHeight = 130.sp,
                style = AppTheme.typography.boldButtonText
            )
        }

        Spacer(modifier = Modifier.height(10.dp))

        Row(
            verticalAlignment = Alignment.CenterVertically,
            modifier = Modifier.fillMaxWidth( ),
            horizontalArrangement = Arrangement.Center
        ) {
            Divider(modifier = Modifier.width(100.dp).background(AppTheme.colors.text), thickness = 1.dp)

            Text(
                text = "or",
                modifier = Modifier.padding(horizontal = 10.dp),
                style = AppTheme.typography.smallSizedBoldText
            )

            Divider(modifier = Modifier.width(100.dp).background(AppTheme.colors.text), thickness = 1.dp)
        }

        Spacer(modifier = Modifier.height(10.dp))

        Button(
            onClick = { },
            shape = RoundedCornerShape(35.dp),
            colors = ButtonDefaults.buttonColors(backgroundColor = Color.White),
            modifier = Modifier.fillMaxWidth( ).height(42.5.dp)
        ) {
            Text(
                text = "Create new account",
                color = Color.Black,
                lineHeight = 130.sp,
                style = AppTheme.typography.boldButtonText
            )
        }
        
        Spacer(modifier = Modifier.height(20.dp))

        Text(
            text = buildAnnotatedString {
                withStyle(style = SpanStyle(
                    color = AppTheme.colors.text)) {append("By signing up, you agree to our ")}
                withStyle(style = SpanStyle(
                    color = AppTheme.colors.primary)) {append("Terms, Privacy Policy ")}
                withStyle(style = SpanStyle(
                    color = AppTheme.colors.text)) {append("and")}
                withStyle(style = SpanStyle(
                    color = AppTheme.colors.primary)) {append(" Cookie Use")}},

            style = AppTheme.typography.smallSizedParagraphText
        )

        Spacer(modifier = Modifier.height(35.dp))

        Text(
            text = buildAnnotatedString {
                withStyle(style = SpanStyle(
                    color = AppTheme.colors.text)) {append("Have an account already? ")}
                withStyle(style = SpanStyle(
                    color = AppTheme.colors.primary)) {append("Sign in ")}},

            style = AppTheme.typography.smallSizedParagraphText
        )
    }
}