package com.twitterclone.theme

import androidx.compose.ui.graphics.Color
import androidx.compose.ui.text.TextStyle
import androidx.compose.ui.text.font.Font
import androidx.compose.ui.text.font.FontFamily
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.sp
import com.twitterclone.R

private val montserratFontFamily= FontFamily(

    Font(R.font.montserrat_regular, FontWeight.Normal),
    Font(R.font.montserrat_medium, FontWeight.Medium),
    Font(R.font.montserrat_semibold, FontWeight.SemiBold),
    Font(R.font.montserrat_bold, FontWeight.Bold),
    Font(R.font.montserrat_extrabold, FontWeight.ExtraBold),
    Font(R.font.montserrat_black, FontWeight.Black)

)

private val graphikFontFamily= FontFamily(

    Font(R.font.graphik_regular, FontWeight.Normal)

)

class ThemeTypography(

    val blackHeading: TextStyle= TextStyle(
        fontFamily = montserratFontFamily,
        fontWeight = FontWeight.Black,
        fontSize = 30.sp,
        lineHeight = 39.sp,
        color = Color.White
    ),

    val extraBoldHeading: TextStyle= TextStyle(
        fontFamily = montserratFontFamily,
        fontWeight = FontWeight.ExtraBold,
        fontSize = 25.sp,
        lineHeight = 32.5.sp,
        color = Color.White
    ),

    val regularSizedBoldText: TextStyle= TextStyle(
        fontFamily = montserratFontFamily,
        fontWeight = FontWeight.Bold,
        fontSize = 16.sp,
        lineHeight = 22.4.sp,
        color = Color.White
    ),

    val smallSizedBoldText: TextStyle= TextStyle(
        fontFamily = montserratFontFamily,
        fontWeight = FontWeight.Bold,
        fontSize = 15.sp,
        lineHeight = 20.25.sp,
        color = Color.White
    ),

    val regularSizedSemiBoldText: TextStyle= TextStyle(
        fontFamily = montserratFontFamily,
        fontWeight = FontWeight.SemiBold,
        fontSize = 15.sp,
        lineHeight = 21.sp,
        color = Color.White
    ),

    val smallSizedSemiBoldText: TextStyle= TextStyle(
        fontFamily = montserratFontFamily,
        fontWeight = FontWeight.SemiBold,
        fontSize = 14.sp,
        lineHeight = 19.6.sp,
        color = Color.White
    ),

    val regularSizedParagraphText: TextStyle= TextStyle(
        fontFamily = montserratFontFamily,
        fontWeight = FontWeight.Medium,
        fontSize = 16.sp,
        lineHeight = 22.4.sp,
        color = Color.White
    ),

    val smallSizedParagraphText: TextStyle= TextStyle(
        fontFamily = montserratFontFamily,
        fontWeight = FontWeight.Medium,
        fontSize = 14.sp,
        lineHeight = 19.6.sp,
        color = Color.White
    ),

    val boldButtonText: TextStyle= TextStyle(
        fontFamily = montserratFontFamily,
        fontWeight = FontWeight.Bold,
        fontSize = 14.sp,
        lineHeight = 19.6.sp
    ),

    val smallSizedPlaceholderText: TextStyle= TextStyle(
        fontFamily = montserratFontFamily,
        fontWeight = FontWeight.Medium,
        fontSize = 13.sp,
        lineHeight = 18.2.sp,
        color = Color.White
    ),

    val largeSizedTweetText: TextStyle= TextStyle(
        fontFamily = graphikFontFamily,
        fontWeight = FontWeight.Normal,
        fontSize = 20.sp,
        lineHeight = 27.sp,
        color = Color.White
    ),

    val regularSizedTweetText: TextStyle= TextStyle(
        fontFamily = graphikFontFamily,
        fontWeight = FontWeight.Normal,
        fontSize = 15.sp,
        lineHeight = 20.25.sp,
        color = Color.White
    )

)